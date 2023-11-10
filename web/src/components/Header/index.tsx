import { Link } from 'react-router-dom';

export function Header() {
  return(
    <div className="header">
      <h1>Cinema Go</h1>

      <ul className="menu">
       <Link to="/admin" className='link'>
          Dashboard
        </Link>
        <Link to="/cadastros" className='link'>
          Cadastros
        </Link>
        <Link to="/config" className='link'>
          Configurações
        </Link>
      </ul>
    </div>
  )
}