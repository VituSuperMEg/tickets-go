import { DashBoard } from "../../components/DashBoard";

export function Admin() {
  return (
    <div className="admin"> 
      <p className="subtitle">
       👋 Seja Bem-vindo, Vitor. 
      </p>
      <div>
        <DashBoard />
      </div>
    </div>
  )
}