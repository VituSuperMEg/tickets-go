export function Login() {
  return (
    <div className="container">
      <form className="login-card">
        <h1>Login<span className="block">.</span></h1>
        <p className="subtitle">Gerencie seu cinema de uma maneira bem diferente<span className="block">!</span></p>
        <div className="login-content">
        <input
          type="text"
          name=""
          id=""
          className="form-control"
          placeholder="login"
        />
        <input
          type="password"
          name=""
          id=""
          className="form-control"
          placeholder="senha"
        />
        <button className="btn-btn-primary">
          Login
        </button>
        <button className="btn-btn-primary-outline">
          Criar Nova Conta
        </button>
        <p className="subtitle">não lembra a senha ? <a className="block" href="#">clique aqui.</a></p>
        </div>
      </form>
    </div>
  );
}
