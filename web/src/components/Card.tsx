import React, { useState } from 'react'
import axios from 'axios'

type Props = {}

const Card = (props: Props) => {
    const [fullName, setfullName] = useState('')
    const [salary, setSalary] = useState(0)
    const handlerSubmit = async (e: React.FormEvent<HTMLFormElement | HTMLButtonElement>) => {
        e.preventDefault()
        try {
            axios.post('http://localhost:30091/api/v1/test', {
                fullName: fullName,
                salary: salary
            })
        } catch (error) {
            console.log(error)
        }
    }


    return (
        <div className="flex items-center justify-center bg-gray-100 min-h-screen">
            <div className="bg-white rounded-lg shadow-lg p-8 w-full max-w-md">
                <h2 className="text-2xl font-semibold mb-6 text-center">Hello</h2>
                <form className='flex flex-col items-center justify-center gap-4'
                    onSubmit={handlerSubmit}
                >
                    <label>FullName</label>
                    <input
                        type='text'
                        name='fullname'
                        className='border rounded-lg shadow-md p-1 min-w-[300px]'
                        onChange={(e) => setfullName(e.target.value)}
                    />
                    <label>Salary</label>
                    <input
                        type='number'
                        name='salary'
                        className='border rounded-lg shadow-md p-1 min-w-[300px]'
                        onChange={(e) => setSalary(Number(e.target.value))}
                    />
                    <button
                        className='bg-blue-500 text-white rounded-lg shadow-md p-2 min-w-[300px]'
                        onClick={handlerSubmit}
                    >
                        Submit
                    </button>
                </form>
            </div>
        </div>
    )
}

export default Card