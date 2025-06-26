import { Children, useEffect } from "react";
import { Footer } from "./elements/Footer";
import { Navbar } from "./elements/Navbar";

interface LayoutProps {
    title: string;
    children: React.ReactNode;
}

const Layout  = ({title,children}: LayoutProps) => {
    useEffect(()=> {
        document.title = title;
    },[title]);

    return <>
        <Navbar/> 
        <main className="flex flex-col gap-y-20 md:gap-y-20 overflow-hidden">{children}</main>
        <Footer/>
    </>
}

export default Layout;
