import { Button, Layout, Menu, message, Row, Col } from "antd";
import { MinusOutlined, PlusSquareOutlined, DeleteOutlined } from "@ant-design/icons";
import { React, useEffect, useRef, useState } from "react";
import Note from "../components/Note";
import { useNavigate } from "react-router-dom";

const { Header, Sider, Content, Footer } = Layout;

async function fetchNotes() {
  const response = await fetch("/api/notes", {
    method: "GET",
  });
  return response;
}

function NotesPage() {
  const navigate = useNavigate();
  const [loadedNotes, setLoadedNotes] = useState([]);
  const [thisNote, setThisNote] = useState({
    Title: "",
    Content: ""
  });

  const menuClickHandler = async (e) => {
    let cnote = loadedNotes.find(x => x.Id == e.key)
    setThisNote({
      Id: cnote.Id,
      Title: cnote.Title,
      Content: cnote.Content
    })
  }

  const deleteHandler = async (Id, e) => {
    console.log("Deleting note ...:", thisNote);
    let response = await fetch("/api/delete", {
      method: "POST",
      body: JSON.stringify({
        id: Id
      }),
    })
    console.log(response);
    if (response.ok) {
      message.success("Note deleted");
      let index = loadedNotes.findIndex(x => x.Id == Id)
      setLoadedNotes([...loadedNotes.slice(0, index), ...loadedNotes.slice(index + 1)])
    } else {
      message.error(response.statusText);
    }
  };

  const saveHandler = async (Id, Title, Content) => {
    if (!Id) return
    console.log("Updating note ...:", thisNote);
    let newNote = {
      Id: Id,
      Title: Title,
      Content: Content
    }
    let response = await fetch("/api/update", {
      method: "POST",
      body: JSON.stringify(newNote),
    })
    console.log(response);
    if (response.ok) {
      message.success("Note updated");
      let index = loadedNotes.findIndex(x => x.Id == Id)
      setLoadedNotes([...loadedNotes.slice(0, index), newNote, ...loadedNotes.slice(index + 1)])
    } else {
      message.error(response.statusText);
    }
  };

  const addHandler = async (e) => {
    console.log("Adding note ...:", thisNote);
    let response = await fetch("/api/add", {
      method: "POST",
      body: JSON.stringify({
        title: 'New Note',
        content: '',
      }),
    })
    console.log(response);
    if (response.ok) {
      message.success("Note added");
      let res = await response.json()
      setLoadedNotes([...loadedNotes, res.newNote])
    } else {
      message.error(response.statusText);
    }
  };

  useEffect(async () => {
    const response = await fetchNotes();
    console.log(response);
    if (!response.ok) {
      message.warning("Please login first");
      navigate("/");
    } else {
      const body = await response.json();
      setLoadedNotes(body.notes);
    }
  }, []);

  return (
    <Layout>
      <Header className="site-layout-background" style={{ padding: 0 }}>
        <Row>
          <h1 style={{ color: "#fff", margin: "0 0.1em 0.5em 0.5em" }}>
            My Notes
          </h1>
          <Button
            onClick={addHandler}
            type="primary"
            style={{ margin: "1.2em 0.5em" }}
          >
            Add <PlusSquareOutlined />
          </Button>
          <div style={{marginLeft: "auto", marginRight: 20}}>
            <a  href="/api/signout">Sign Out</a>
          </div>
        </Row>
      </Header>
      <Layout hasSider className="site-layout">
        <Sider>
          <Menu theme="dark" mode="inline" defaultSelectedKeys={["1"]} className="mscroll">
            {loadedNotes.map((note) => (
              <Menu.Item key={note.Id} icon={<MinusOutlined />} onClick={menuClickHandler} className="miscroll">
                {note.Title}
                <DeleteOutlined style={{marginTop: 12, float: 'right'}} onClick={deleteHandler.bind(null, note.Id)} />
              </Menu.Item>
            ))}
          </Menu>
        </Sider>
        <Content
          className="site-layout-background"
          style={{
            margin: "24px 16px",
            padding: 24,
            minHeight: "100vh",
          }}
        >
          <Note
            save={saveHandler}
            id={thisNote.Id}
            title={thisNote.Title}
            content={thisNote.Content}
          />
        </Content>
      </Layout>
    </Layout>
  );
}

export default NotesPage;
