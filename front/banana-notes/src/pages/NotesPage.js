import { Button, Layout, Menu, message, Row } from "antd";
import { MinusOutlined, PlusSquareOutlined } from "@ant-design/icons";
import { React, useEffect, useRef, useState } from "react";
import Note from "../components/Note";
import { useNavigate } from "react-router-dom";

const { Header, Sider, Content, Footer } = Layout;

async function isLoggedIn() {
  const response = await fetch("/api/user", {
    method: "GET",
  });
  return response.ok;
}

async function fetchNotes() {
  const response = await fetch("/api/notes", {
    method: "GET",
  });
  return response;
}

function NotesPage() {
  const navigate = useNavigate();
  const [loadedNotes, setLoadedNotes] = useState([]);
  const currentNote = useRef(null);

  const addHandler = (e) => {
    console.log("Adding note ...:", currentNote);
    fetch("/api/add", {
      method: "POST",
      body: JSON.stringify({
        title: 'Fake Title',
        content: 'Faker Content',
      }),
    }).then((response) => {
      console.log(response);
      if (response.ok) {
        message.success("Note added");
      } else {
        message.error(response.statusText);
      }
    });
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
        </Row>
      </Header>
      <Layout hasSider className="site-layout">
        <Sider>
          <Menu theme="dark" mode="inline" defaultSelectedKeys={["1"]}>
            {loadedNotes.map((note) => (
              <Menu.Item key={note.Id} icon={<MinusOutlined />}>
                {note.Title}
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
            ref={currentNote}
            title="NEW NOTE"
            content="Enter your note here"
          />
        </Content>
      </Layout>
    </Layout>
  );
}

export default NotesPage;
