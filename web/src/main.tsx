import React from 'react';
import ReactDOM from 'react-dom';
import { App } from './app';
import './styles/ready.css'
const Main = () => {
  return (
    <App />
  );
};

ReactDOM.render(<Main />, document.getElementById('root'));