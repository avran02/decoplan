'use client'
import { refreshInterceptor } from '@/api/authRefreshToken'
import { getAccessToken } from '@/services/auth.helper'
import axios, { CreateAxiosDefaults } from 'axios'
import { useState } from 'react'

const axiosOptions: CreateAxiosDefaults = {
    baseURL: "http://files.localhost",
    withCredentials: true,
}

const CheckFileExists = () => {
    const instance = axios.create(axiosOptions)
    refreshInterceptor(instance)
    const [fileId, setFileId] = useState('')
    const [exists, setExists] = useState(null)
    const [error, setError] = useState('')

    const handleCheckFileExists = async () => {
        try {
            const response = await instance.get(`exists/${fileId}`, {
                headers: {
                    'Authorization': `Bearer ${getAccessToken()}`
                }
            })
            setExists(response.data.ok)
            setError('')
        } catch (error) {
            setExists(null)
            setError('Ошибка: Невозможно проверить наличие файла')
        }
    }

    return (
        <div
            background-color='#'>
            <h1>Check if File Exists</h1>
            <input
                type="text"
                style={{ color: '#333333' }}
                placeholder="Enter file ID"
                value={fileId}
                onChange={(e) => setFileId(e.target.value)}
            />
            <button onClick={handleCheckFileExists}>Check File Exists</button>

            {error && <p>{error}</p>}
            {exists !== null && <p>File Exists: {exists ? 'Yes' : 'No'}</p>}
        </div>
    )
}

export default CheckFileExists
