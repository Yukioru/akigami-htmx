import { Config } from "tailwindcss";

const config: Config = {
  content: ["./views/**/*.jet"],
  theme: {
    extend: {
      container: {
        center: true,
        padding: "2rem",
      },
    },
  },
  plugins: [],
};

export default config;
