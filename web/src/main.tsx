import { createRoot } from 'react-dom/client';
import App from '@/App';
import { CssVarsProvider } from '@mui/joy';
import { ToastContainer } from 'react-toastify'; //toastify

import 'react-toastify/dist/ReactToastify.css'; //toastify css
import theme from './styles/theme';
import './styles/global.css';
import './styles/tailwind.css';

createRoot(document.getElementById('root')!).render(
  <CssVarsProvider theme={theme}>
    <App />
    <ToastContainer />
  </CssVarsProvider>
);
