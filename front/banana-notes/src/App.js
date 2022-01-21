import { Route, Routes } from "react-router-dom";

import "./App.css";
import EntryPage from "./pages/EntryPage";

function App() {
  return (
    <Routes>
      <Route path="/" element={<EntryPage />} />
      <Route path="/notes" />
    </Routes>
  );
}

export default App;
