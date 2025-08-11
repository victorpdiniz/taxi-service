import { defineConfig } from "cypress";
import { spawn } from "child_process"; // Changed from exec to spawn

export default defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      on('task', {
        runDriverSimulator({ rideId, scenario }) {
          return new Promise((resolve, reject) => {
            const command = 'node';
            const args = ['driver_simulator.js', `--corrida-id=${rideId}`, `--cenario=${scenario}`];
            const options = { cwd: '../' }; // Run from project root

            const child = spawn(command, args, options);

            let output = '';
            child.stdout.on('data', (data) => {
              output += data.toString();
              console.log(`driver_simulator stdout: ${data}`);
              if (output.includes('CYPRESS_TASK_RIDE_ACCEPTED')) {
                resolve(true); // Resolve the task as soon as the signal is received
              }
            });

            child.stderr.on('data', (data) => {
              console.error(`driver_simulator stderr: ${data}`);
            });

            child.on('close', (code) => {
              if (code !== 0 && !output.includes('CYPRESS_TASK_RIDE_ACCEPTED')) {
                reject(new Error(`driver_simulator exited with code ${code}`));
              }
            });

            child.on('error', (err) => {
              reject(err);
            });
          });
        },
      });
      return config;
    },
    baseUrl: 'http://127.0.0.1:5173',
  },
});