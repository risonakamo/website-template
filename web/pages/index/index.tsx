import {createRoot} from "react-dom/client";

import "./index.less";

function IndexMain():JSX.Element
{
  return <>
    huh
  </>;
}

function main()
{
  createRoot(document.querySelector(".main")!).render(<IndexMain/>);
}

window.onload=main;