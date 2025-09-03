import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0', // Permite conexiones desde cualquier IP
    port: 5173,
    strictPort: true, // Falla si el puerto no está disponible
    watch: {
      usePolling: true, // Necesario para hot reload en Docker
      interval: 1000,   // Intervalo de polling en ms
    },
  },
  build: {
    outDir: 'dist',
  },
})