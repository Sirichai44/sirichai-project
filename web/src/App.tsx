import { useEffect, useState } from 'react';
import { Outlet } from 'react-router';
import { Button, useColorScheme } from '@mui/joy';
function App() {
  const { mode, setMode } = useColorScheme();
  const [loading, setLoading] = useState(true);
  useEffect(() => {
    const initialState = async () => {};

    Promise.all([initialState()]).then(() => setLoading(false));
  }, []);

  useEffect(() => {
    const root = document.documentElement;

    if (mode === 'light') {
      root.classList.remove('dark');
    } else if (mode === 'dark') {
      root.classList.add('dark');
    }
  }, [mode]);

  // return loading ? null : <Outlet />;
  return (
    <div>
      <Button
        variant="outlined"
        onClick={() => {
          setMode(mode === 'light' ? 'dark' : 'light');
        }}>
        Hello World
      </Button>
      <Outlet />
    </div>
  );
}

export default App;
