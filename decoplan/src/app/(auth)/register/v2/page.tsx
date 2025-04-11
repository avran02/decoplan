import type { Metadata } from "next"
import { AuthFormV2 } from "./AuthFormV2"

export const metadata: Metadata = {
	title: "",
	robots: {
		index: false,
		follow: false,
	},
}

export default function RegistationV2Page() {
	return (
		<div className='min-h-screen flex items-center justify-center bg-neutral-950 font-[family-name:var(--font-monocraft)]'>
			<div className='bg-neutral-900 p-8 rounded-lg shadow-md'>
				<h2 className='font-bold mb-4'>Ещё немного✌️</h2>
				<AuthFormV2 />
			</div>
		</div>
	)
}
