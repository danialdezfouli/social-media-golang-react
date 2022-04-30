import Layout from "components/layout/Layout";
import Home from "pages/Home";
import Login from "pages/Login";
import NotFound from "pages/NotFound";
import PostPage from "pages/PostPage";
import Profile from "pages/Profile";
import Register from "pages/Register";
import Search from "pages/Search";
import Settings from "pages/Settings";
import { Route, Routes } from "react-router-dom";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Login />} />
      <Route path="signup" element={<Register />} />
      <Route path="/" element={<Layout />}>
        <Route path="home" element={<Home />} />
        <Route path="search" element={<Search />} />
        <Route path="profile/:id" element={<Profile />} />
        <Route path="post/:id" element={<PostPage />} />
        <Route path="settings" element={<Settings />} />
        <Route path="*" element={<NotFound />} />
      </Route>
    </Routes>
  );
}

export default App;
