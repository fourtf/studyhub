import './App.css';
import Quizzes from './pages/Quizzes';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Home from 'pages/Home';
import NotFound from 'pages/NotFound';
import Quiz from 'pages/Quiz';

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path="/" exact component={Home} />
          <Route path="/quizzes" exact component={Quizzes} />
          <Route path="/quizzes/:id" exact component={Quiz} />
          <Route component={NotFound} />
        </Switch>
      </Router>
    </div>
  );
}

export default App;
