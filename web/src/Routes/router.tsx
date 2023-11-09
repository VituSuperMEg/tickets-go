import { Routes, Route } from "react-router-dom";
import { Login } from "../shared/login";
import { Admin } from "../modules/Admin";

export function Router() {
  return (
    <Routes>
      <Route path="/" element={<Login />}/>
      <Route path="/admin" element={<Admin />} />
    </Routes>
  )
}