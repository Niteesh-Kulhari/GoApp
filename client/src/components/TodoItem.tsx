import { Badge, Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import type { Todo } from "./TodoList";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export default function TodoItem({todo} : {todo: Todo}) {
    const queryClient = useQueryClient();
    const {mutate: updateTodo, isPending: isUpdating} =useMutation({
		mutationKey: ["updateTodo"],
		mutationFn: async () => {
			if(todo.completed) return alert("Todo is already Completed")

                try {
                    const res = await fetch(`http://localhost:4000/api/todos/${todo._id}`, {
                        method: "PATCH",
                    })
                    const data = await res.json()
                    if(!res.ok){
                        throw new Error(data.error || "Something went wrong")
                    }
                    return data
                } catch (error) {
                    console.log(error);
                }
		},
        onSuccess: () => {
            queryClient.invalidateQueries({queryKey: ["todos"]})
        }
	}) 

    return (
		<Flex gap={2} alignItems={"center"}>
			<Flex
				flex={1}
				alignItems={"center"}
				border={"1px"}
				borderColor={"gray.600"}
				p={2}
				borderRadius={"lg"}
				justifyContent={"space-between"}
			>
				<Text
					color={todo.completed ? "green.200" : "yellow.100"}
					textDecoration={todo.completed ? "line-through" : "none"}
				>
					{todo.body}
				</Text>
				{todo.completed && (
					<Badge ml='1' colorScheme='green'>
						Done
					</Badge>
				)}
				{!todo.completed && (
					<Badge ml='1' colorScheme='yellow'>
						In Progress
					</Badge>
				)}
			</Flex>
			<Flex gap={2} alignItems={"center"}>
				<Box color={"green.500"} cursor={"pointer"} onClick={() => updateTodo()}>
					{!isUpdating && <FaCheckCircle size={20} /> }
                    {isUpdating && <Spinner/>}
				</Box>
				<Box color={"red.500"} cursor={"pointer"}>
					<MdDelete size={25} />
				</Box>
			</Flex>
		</Flex>
	);
}