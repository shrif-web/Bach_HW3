import { Card, Typography } from "antd";
import { useState } from "react";
const { Paragraph } = Typography;

function Note(props) {
  const [title, setTitle] = useState(props.title);
  const [content, setContent] = useState(props.content);

  const onTitleUpdate = (e) => {
    if (e.length == 0) {
      e = "Title";
    }
    setTitle(e);
  };
  const onContentUpdate = (e) => {setContent(e)};

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
