import {createBrowserRouter, createRoutesFromElements, Route, RouterProvider} from "react-router-dom";
import React, {useEffect} from "react";
import Layout from "./components/layouts/Layout";
import LoginPage from "./routes/LoginPage";
import ProtectedLayout from "./components/layouts/ProtectedLayout";
import GroupPage from "./routes/GroupPage";
import useAuthentication from "./stores/useAuthentication";
import HomePage from "./routes/HomePage";
import ErrorPage from "./routes/ErrorPage";
import RegisterPage from "./routes/RegisterPage";

const router = createBrowserRouter(
    createRoutesFromElements(
        <>
            <Route path="/" element={<ProtectedLayout/>} errorElement={<ErrorPage/>}>
                <Route path="/" element={<HomePage/>}/>
                <Route path="/groups/:id" element={<GroupPage/>}/>
            </Route>

            <Route path="/" element={<Layout/>} errorElement={<ErrorPage/>}>
                <Route path="/login" element={<LoginPage/>}/>
                <Route path="/register" element={<RegisterPage/>}/>
            </Route>
        </>
    )
);
const App = () => {
    const {init: preferencesInit} = useAuthentication();
    const {init: authInit} = useAuthentication();

    useEffect(() => {
        preferencesInit();
        authInit();
    }, [authInit]);

    return (
        <RouterProvider router={router}/>
    )
}

export default App
