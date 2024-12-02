import { useEffect, useState } from 'react';
import { Outlet } from 'react-router';
import { useColorScheme } from '@mui/joy';
function App() {
  const { mode } = useColorScheme();
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

  return loading ? null : <Outlet />;
  // return <h1 className="text-3xl font-bold underline">Hello world!</h1>;
}

export default App;
