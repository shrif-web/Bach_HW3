import { Card, Typography } from "antd";
import { useState, useEffect } from "react";
const { Paragraph } = Typography;

function Note(props) {
  const [title, setTitle] = useState(props.title);
  const [content, setContent] = useState(props.content);

  useEffect(() => {
    setTitle(props.title);
  }, [props.title])

  useEffect(() => {
    setContent(props.content);
  }, [props.content])

  const onTitleUpdate = (e) => {
    if (e.length == 0) {
      e = "Title";
    }
    setTitle(e);
    props.save(props.id, e, content)
  };
  const onContentUpdate = (e) => {
    setContent(e)
    props.save(props.id, title, e)
  };

  return (
    <Card
      title={
        <Paragraph
          editable={{
            onChange: onTitleUpdate,
            tooltip: "Click to edit text",
            triggerType: "[text]",
          }}
        >
          {title}
        </Paragraph>
      }
    >
      <Paragraph
        editable={{
          onChange: onContentUpdate,
          tooltip: "Click to edit text",
          triggerType: "[text, icon]",
        }}
      >
        {content}
      </Paragraph>
    </Card>
  );
}

export default Note;
