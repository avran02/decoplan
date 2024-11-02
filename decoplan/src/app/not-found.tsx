import { NotFound } from "@/components/ui/not-found/NotFound"
import type { Metadata } from "next"

export const metadata: Metadata = {
	title: "",
}

export default function NotFoundPage() {
	return <NotFound />
}
