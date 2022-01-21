import { Button, Form, Input } from "antd";

function SignInForm(props) {
  return (
    <Form
      name="signIn"
      onFinish={props.onFinish}
      onFinishFailed={props.onFinishFailed}
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
        <Button style={{ marginTop: "10px" }} onClick={props.goToSignUp} block>
          Sign Up
        </Button>
      </Form.Item>
    </Form>
  );
}

export default SignInForm