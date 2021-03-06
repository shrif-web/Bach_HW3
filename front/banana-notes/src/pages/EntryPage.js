import { Card, Typography, message } from "antd";
import { useEffect, useState } from "react";
import { useNavigate } from 'react-router-dom'

import SignInForm from "../components/SignInForm";
import SignUpForm from "../components/SignUpForm";

const { Meta } = Card;
const { Title } = Typography;

function EntryPage() {
  const [state, updateState] = useState("signin");
  const navigate = useNavigate();

  async function isLoggedIn() {
    const response = await fetch("/api/user", {
      method: "GET",
    });
    return response.ok;
  }

  useEffect(async () => {
    if (await isLoggedIn()) {
      navigate("/notes")
    }
  }, [])

  const onSignInFinish = (values) => {
    console.log("Signing in ...:", values);
    fetch("/api/auth", {
      method: "POST",
      body: JSON.stringify({
        user: values.username,
        pass: values.password,
      }),
    }).then((response) => {
      console.log(response);
      if (response.ok) {
        navigate('/notes')
      } else {
        message.error(response.statusText);
      }
    });
  };

  const onSignInFinishFailed = (errorInfo) => {
    console.log("Failed:", errorInfo);
  };

  const goToSignUp = () => {
    console.log('go to sign up')
    updateState("signup");
  };

  const goToSignIn = () => {
    console.log('go to sign in')
    updateState("signin");
  };

  const onSignUpFinish = (values) => {
    console.log("Signing up ...:", values);
    fetch("/api/reg", {
      method: "POST",
      body: JSON.stringify({
        firstname: values.firstname,
        lastname: values.lastname,
        user: values.username,
        pass: values.password,
      }),
    }).then((response) => {
      console.log(response);
      if (response.ok) {
        navigate('/notes')
      } else {
        message.error(response.statusText);
      }
    });
  };

  const onSignUpFinishFailed = (errorInfo) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <center>
      <Card
        hoverable
        style={{ width: 400, marginTop: "2em" }}
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
        {state == "signin" && (
          <SignInForm
            onFinish={onSignInFinish}
            onFinishFailed={onSignInFinishFailed}
            goToSignUp={goToSignUp}
          />
        )}
        {state == "signup" && (
          <SignUpForm
            onFinish={onSignUpFinish}
            onFinishFailed={onSignUpFinishFailed}
            goToSignIn={goToSignIn}
          />
        )}
      </Card>
    </center>
  );
}

export default EntryPage;
