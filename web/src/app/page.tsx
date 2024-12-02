import { Button, useColorScheme } from '@mui/joy';
import { Suspense } from 'react';
import { Outlet } from 'react-router';

const Page = () => {
  const { mode, setMode } = useColorScheme();
  return (
    <div className="flex w-full h-screen border border-red-500">
      <div className="w-32 border border-red-500">
        {' '}
        <Button
          variant="outlined"
          className="bg-red-500"
          onClick={() => {
            setMode(mode === 'light' ? 'dark' : 'light');
          }}>
          Hello World
        </Button>
      </div>

      <main className="flex flex-col items-center justify-center flex-grow w-full h-auto shrink">
        <Suspense>
          <Outlet />
        </Suspense>
      </main>
    </div>
  );
};

export default Page;
