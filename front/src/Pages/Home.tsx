import React from 'react';

type Task = {
  id: number;
  title: string;
};

type HomeProps = {
  name: string;
  tasks: Task[];
};

const Home: React.FC<HomeProps> = ({ name, tasks }) => {
  console.log("Renderizando componente Home", { name, tasks });
  return (
    <div>
      <h1>{name}</h1>
      <ul>
        {tasks.map((task) => (
          <li key={task.id}>{task.title}</li>
        ))}
      </ul>
    </div>
  );
};

export default Home;

