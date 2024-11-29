import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';
import path from 'path';

const frontPort = 3000;
const serverHost = 'http://localhost:8080';

export default defineConfig({
  plugins: [react()],
  server: {
    port: frontPort,
    proxy: {
      '/api': {
        target: serverHost,
        changeOrigin: true
      }
    }
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@components': path.resolve(__dirname, './src/components')
    }
  }
});
