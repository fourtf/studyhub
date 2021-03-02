import React from 'react';
import Nav from 'components/Nav';
import { useParams } from 'react-router-dom';

function Quiz() {
  let { id } = useParams<{ id: string }>();

  console.log(id);
  return (
    <div>
      <Nav />
      Quiz: {id}
    </div>
  );
}

export default Quiz;
