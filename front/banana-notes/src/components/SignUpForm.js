import { Button, Form, Input, Typography} from "antd";
const { Text, Link } = Typography;


function SignUpForm(props) {
  return (
    <Form
      name="signIn"
      onFinish={props.onFinish}
      onFinishFailed={props.onFinishFailed}
      autoComplete="off"
    >
      <Form.Item
        label="First Name"
        name="firstname"
        rules={[
          {
            required: true,
            message: "Please input your first name!",
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Last Name"
        name="lastname"
        rules={[
          {
            required: true,
            message: "Please input your last name!",
          },
        ]}
      >
        <Input />
      </Form.Item>

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
          Sign Up
        </Button>
        <div style={{ marginTop: "10px" }} >
          Or <Button onClick={props.goToSignIn} type='text' style={{padding: 0}}><Text underline>log in</Text></Button> instead
        </div>
      </Form.Item>
    </Form>
  );
}

export default SignUpForm
