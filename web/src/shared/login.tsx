import { useFormik } from "formik";
import { useState } from "react";
import { api } from "../services/api";
import Swal from "sweetalert2";
import { useNavigate } from "react-router-dom";

export function Login() {
  const [isValid, setIsValid] = useState(false);
  const navigate = useNavigate()
  const isLogin = useFormik({
    initialValues : {
      login: "",
      password: "",
    },
    onSubmit : async (values) => {
      try {
        const response = await api.post("/login", {
          login: values.login,
          password: values.password,
        });
        console.log(response.data);
        if (response.data.token) {
          navigate("/admin");
        }
      } catch (error) {
        if (error.response && error.response.data && error.response.data.error) {
          Swal.fire({
            text: error.response.data.error,
            icon: "error",
          });
        }
      }
    }
  })
  return (
    <div className="container">
      {!isValid ? (
        <form className="login-card" onSubmit={isLogin.handleSubmit}>
          <h1>
            Login<span className="block">.</span>
          </h1>
          <p className="subtitle">
            Gerencie seu cinema de uma maneira bem diferente
            <span className="block">!</span>
          </p>
          <div className="login-content">
            <input
              type="text"
              name="login"
              id="login"
              className="form-control"
              placeholder="login"
              onChange={isLogin.handleChange}
               value={isLogin.values.login}
            />
            <input
              type="password"
              name="password"
              id="password"
              className="form-control"
              placeholder="senha"
              onChange={isLogin.handleChange}
               value={isLogin.values.password}
            />
            <button className="btn-btn-primary" type="submit">Login</button>
            <button className="btn-btn-primary-outline">
              Criar Nova Conta
            </button>
            <p className="subtitle">
              não lembra a senha ?{" "}
              <a className="block" href="#">
                clique aqui.
              </a>
            </p>
          </div>
        </form>
      ) : (
        "Login  is successfully"
      )}
    </div>
  );
}
