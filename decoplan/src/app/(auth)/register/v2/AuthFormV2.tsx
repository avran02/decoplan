"use client"

import clsx from "clsx"
import { useRouter } from "next/navigation"

import { Date } from "@/components/ui/DatePicker/Date"
import usersChatsService from "@/services/usersChats.service"
import { IFormRegisterV2 } from "@/types/usersChats.types"
import { useMutation } from "@tanstack/react-query"
import { useForm } from "react-hook-form"
import styles from "../../AuthForm.module.scss"

export function AuthFormV2() {
	const { register, control, handleSubmit, reset } = useForm()
	const router = useRouter()

	const { mutate: mutateRegister } = useMutation({
		mutationKey: ["register"],
		mutationFn: (data: IFormRegisterV2) => usersChatsService.createUser(data),
		onSuccess() {
			reset()
			router.push("/chat")
		},
	})

	const onSubmit = (data: any) => {
		mutateRegister(data)
	}

	return (
		<form
			className='font-[family-name:var(--font-monocraft)]'
			onSubmit={handleSubmit(onSubmit)}
		>
			<div className='mb-4'>
				<label className='text-gray-600'>
					Username
					<input
						type='username'
						placeholder='Enter login: '
						{...register("username", { required: true })}
						className={clsx(
							styles["input-field"],
							"w-full p-2 border rounded focus:outline-none focus:border-indigo-500"
						)}
					/>
				</label>
				<label className='text-gray-600'>
					Date
					<Date control={control} defaultValue={null} />
				</label>
			</div>
			<div className='mb-4 col-span-2'>
				<button
					type='submit'
					className={clsx(styles["btn-primary"], "bg-indigo-500")}
				>
					Create
				</button>
			</div>
		</form>
	)
}
