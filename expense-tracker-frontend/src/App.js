import './App.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './components/Login';
import Register from './components/Register';

function App() {
  return (
    <Router>
    <div>
      Hello World
        <Routes>
            <Route path="/login" element={<Login/>} />
            <Route path="/register" element={<Register/>} />
            <Route path="/" exact component={() => <h2>Welcome to Expense Tracker</h2>} />
        </Routes>
    </div>
</Router>
  );
}

export default App;
