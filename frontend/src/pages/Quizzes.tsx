import { Card } from '@material-ui/core';
import Nav from 'components/Nav';
import { Link } from 'react-router-dom';

function Quiz() {
  return (
    <Link to="quizzes/123">
      <Card style={{ width: 300, height: 150, zIndex: 3 }}>sadf</Card>
    </Link>
  );
}

function Quizzes() {
  return (
    <div>
      <Nav />
      <h1>Category</h1>
      <Quiz />
    </div>
  );
}

export default Quizzes;
