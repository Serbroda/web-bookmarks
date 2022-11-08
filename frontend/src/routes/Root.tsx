import { Link, Outlet, useLoaderData } from "react-router-dom";

export interface GroupDto {
  id: number;
  name: string;
}

export const loader = async () => {
  return {
    groups: [
      { id: 1, name: "Rezepte" },
      { id: 2, name: "Videos" },
    ],
  };
};

const Root = () => {
  const { groups } = useLoaderData() as { groups: GroupDto[] };
  console.log(groups);

  return (
    <div className="App">
      <h1>Hi</h1>
      <nav>
        <ul>
          <li>
            <Link to={`/`}>Home</Link>
          </li>
          {groups.map((group) => (
            <li key={group.id}>
              <Link to={`groups/${group.id}`}>{group.name}</Link>
            </li>
          ))}
        </ul>
      </nav>
      <Outlet />
    </div>
  );
};

export default Root;
