import { LinearProgress } from '@mui/joy';
import { Suspense, lazy } from 'react';
import { createBrowserRouter } from 'react-router';

const App = lazy(() => import('@/App'));
const Home = lazy(() => import('@/app/page'));

const router = createBrowserRouter([
  {
    path: '/',
    element: (
      <Suspense fallback={<LinearProgress />}>
        <App />
      </Suspense>
    ),
    children: [{ path: '/', element: <Home /> }]
  }
]);

export default router;
