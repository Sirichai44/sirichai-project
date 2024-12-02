import { createRoot } from 'react-dom/client';
import { RouterProvider } from 'react-router';
import { CssVarsProvider } from '@mui/joy';
import { ToastContainer } from 'react-toastify'; //toastify

import router from '@/route/index';

import 'react-toastify/dist/ReactToastify.css'; //toastify css
import theme from './styles/theme';
import './styles/tailwind.css';
import './styles/global.css';

createRoot(document.getElementById('root')!).render(
  <CssVarsProvider theme={theme}>
    <RouterProvider router={router} />
    <ToastContainer />
  </CssVarsProvider>
);
