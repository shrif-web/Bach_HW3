import { Card, Typography } from "antd";

import { SignInForm } from "../components/SignInForm";

const { Meta } = Card;
const { Title } = Typography;

function EntryPage() {
  return (
    <center>
      <div margin-top="100pt">
        <Card
          hoverable
          style={{ width: 400 }}
          cover={
            <Title style={{ marginTop: "1em", marginBottom: 0 }}>
              Banana Notes
            </Title>
          }
        >
          <Meta
            style={{ marginBottom: "1em" }}
            description="A place to store your notes"
          />
          <SignInForm />
        </Card>
      </div>
    </center>
  );
}

export default EntryPage;
