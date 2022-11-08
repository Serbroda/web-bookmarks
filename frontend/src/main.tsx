import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider, Route } from "react-router-dom";
import "./index.css";
import ErrorPage from "./routes/ErrorPage";
import GroupPage from "./routes/GroupPage";
import HomePage from "./routes/HomePage";
import Root, { loader as rootLoader } from "./routes/Root";
import "tippy.js/dist/tippy.css";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <ErrorPage />,
    loader: rootLoader,
    children: [
      {
        path: "",
        element: <HomePage />,
      },
      {
        path: "groups/:id",
        element: <GroupPage />,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
