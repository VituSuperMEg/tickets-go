import { BrowserRouter } from 'react-router-dom';
import { Router } from './Routes/router';

export function App() {
  return (
    <BrowserRouter>
      <Router />
    </BrowserRouter>
  )
}