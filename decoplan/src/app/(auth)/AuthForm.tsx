"use client"

import authService from "@/services/auth.service"
import { useMutation } from "@tanstack/react-query"
import clsx from "clsx"
import { useRouter } from "next/navigation"
import { SubmitHandler, useForm } from "react-hook-form"
import toast from "react-hot-toast"

import { IFormLogin, IFormRegister } from "@/types/auth.types"

import Link from "next/link"
import styles from "./AuthForm.module.scss"

interface AuthFormProps {
	isLogin: boolean
}

export function AuthForm({ isLogin }: AuthFormProps) {
	const { register, handleSubmit, reset } = useForm<
		IFormLogin | IFormRegister
	>()

	const router = useRouter()

	const { mutate: mutateLogin, isPending: isLoginPending } = useMutation({
		mutationKey: ["login"],
		mutationFn: (data: IFormLogin) => authService.login(data),
		onSuccess() {
			reset()
			router.push("/chat")
			toast.success("Login successful")
		},
	})

	const {
		error,
		mutate: mutateRegister,
		isPending: isRegisterPending,
	} = useMutation({
		mutationKey: ["register"],
		mutationFn: (data: IFormRegister) => authService.register(data),
		onError() {
			// Выводим ошибку в консоль
			console.error(error)
			toast.error("Registration failed")
		},
		onSuccess() {
			reset()
			router.push("/register/v2")
			toast.success("Registration successful")
		},
	})

	const isPending = isLoginPending || isRegisterPending

	const onSubmit: SubmitHandler<IFormLogin | IFormRegister> = (data) => {
		if (isLogin) {
			const formData = data as IFormLogin
			mutateLogin(formData)
		} else {
			const formData = data as IFormRegister
			mutateRegister(formData)
		}
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
						placeholder='Enter username: '
						{...register("username", { required: true })}
						className={clsx(
							styles["input-field"],
							"w-full p-2 border rounded focus:outline-none focus:border-indigo-500"
						)}
					/>
				</label>
			</div>
			{isLogin ? null : (
				<div className='mb-4 col-span-2'>
					<label className='text-gray-600'>
						Email
						<input
							type='email'
							placeholder='Enter email: '
							{...register("email", { required: true })}
							className={clsx(
								styles["input-field"],
								"w-full p-2 border rounded focus:outline-none focus:border-indigo-500"
							)}
						/>
					</label>
				</div>
			)}

			<div className='mb-4 col-span-2'>
				<label className='text-gray-600'>
					Password
					<input
						type='password'
						placeholder='Enter password: '
						{...register("password", { required: true })}
						className={clsx(
							styles["input-field"],
							"w-full p-2 border rounded focus:outline-none focus:border-indigo-500"
						)}
					/>
				</label>
			</div>
			{isLogin ? (
				<Link className={styles.link} href='/register'>
					Create an account
				</Link>
			) : (
				<Link className={styles.link} href='/login'>
					Already have an account? Sign in{" "}
				</Link>
			)}
			<div className='mb-4 col-span-2'>
				<button
					type='submit'
					className={clsx(
						styles["btn-primary"],
						isLogin ? "bg-indigo-500" : "bg-indigo-500",
						isPending ? "opacity-75 cursor-not-allowed" : ""
					)}
					disabled={isPending}
				>
					{isLogin ? "Sign in" : "Sign up"}
				</button>
			</div>
		</form>
	)
}
