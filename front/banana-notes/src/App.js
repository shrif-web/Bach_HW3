import { Button, Result } from "antd";
import { Route, Routes, useNavigate } from "react-router-dom";

import "./App.css";
import EntryPage from "./pages/EntryPage";
import NotesPage from "./pages/NotesPage";

function App() {
  return (
    <Routes>
      <Route path="*" element={<NotFoundPage />} />
      <Route path="/" element={<EntryPage />} />
      <Route path="/notes" element={<NotesPage />} />
    </Routes>
  );
}

function NotFoundPage() {
  const navigate = useNavigate();

  return (
    <Result
      status="404"
      title="404"
      subTitle="Sorry, the page you visited does not exist."
      extra={<Button onClick={()=>navigate('/notes')} type="primary">Back Home</Button>}
    />
  );
}

export default App;
