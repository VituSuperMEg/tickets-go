import { Routes, Route } from "react-router-dom";
import { Login } from "../shared/login";
import { Admin } from "../modules/Admin";
import { DefaultLayout } from "../_layout/AdminLayout";

export function Router() {
  return (
    <Routes>
      <Route path="/" element={<Login />}/>
      <Route path="/admin" element={<DefaultLayout />}>
        <Route path="/admin" element={<Admin />} />
      </Route>
    </Routes>
  )
}