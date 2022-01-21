import { Button, Form, Input, message } from "antd";

export function SignInForm() {

  const onFinish = (values) => {
    console.log("Signing in ...:", values);
    fetch("/auth", {
      method: "POST",
      body: JSON.stringify({
        user: values.username,
        pass: values.password,
      }),
    }).then((response) => {
      console.log(response);
      if (response.ok) {
        //todo Route to Notes Page if successful else show error message
      } else {
        message.error(response.statusText);
      }
    });
  };

  const onFinishFailed = (errorInfo) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <Form
      name="signIn"
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      <Form.Item
        label="Username"
        name="username"
        rules={[
          {
            required: true,
            message: "Please input your username!",
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Password"
        name="password"
        rules={[
          {
            required: true,
            message: "Please input your password!",
          },
        ]}
      >
        <Input.Password />
      </Form.Item>

      <Form.Item>
        <Button type="primary" htmlType="submit" block>
          Sign In
        </Button>
        <Button style={{ marginTop: "10px" }} block>
          Sign Up
        </Button>
      </Form.Item>
    </Form>
  );
}
