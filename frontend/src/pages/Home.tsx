import Nav from 'components/Nav';
import { Link } from 'react-router-dom';

function Home() {
  return (
    <div>
      <Nav />
      <Link to="/quizzes">Quizzes</Link>
    </div>
  );
}

export default Home;
