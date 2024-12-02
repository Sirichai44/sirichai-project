import { Suspense } from 'react';
import { Outlet } from 'react-router';
import { LinearProgress } from '@mui/joy';

function SuspenseWrapper() {
  return (
    <Suspense fallback={<LinearProgress />}>
      <Outlet />
    </Suspense>
  );
}

export default SuspenseWrapper;
