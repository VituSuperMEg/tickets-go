import { Routes, Route } from "react-router-dom";
import { Login } from "../shared/login";
import { Admin } from "../modules/Admin";
import { DefaultLayout } from "../_layout/AdminLayout";
import { Session } from "../components/CardFilme/Session";
import { Cadastros } from "../shared/pages/Cadastros";
import Config from "../shared/pages/Config";

export function Router() {
  return (
    <Routes>
      <Route path="/" element={<Login />}/>
      <Route path="/" element={<DefaultLayout />}>
        <Route path="/admin" element={<Admin />} />
        <Route path="/lotacao/:id" element={<Session />} />
        <Route path="/cadastros" element={<Cadastros />} />
        <Route path="/config" element={<Config />} />
      </Route>
    </Routes>
  )
}