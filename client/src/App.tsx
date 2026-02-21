import { Container, Stack } from "@chakra-ui/react";
import Navbar from "./components/Navbar";
import TodoForm from "./components/TodoForm";
import TodoList from "./components/TodoList";

export const BASE_URL = import.meta.env.MODE === "development" ? "https://rodqg2r5k7.execute-api.ap-south-1.amazonaws.com/api" : "/api";

function App() {
	return (
		<Stack h='100vh'>
			<Navbar />
			<Container>
				<TodoForm />
				<TodoList />
			</Container>
		</Stack>
	);
}

export default App;
