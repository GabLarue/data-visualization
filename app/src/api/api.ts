import axios, { AxiosInstance } from "axios";

class Api {
    createAxios = (contentType = "application/json"): AxiosInstance => {
        return axios.create({
            baseURL: '//localhost:8080',
            headers: { contentType }
        });
    };

    getAllFiles = async () => {
        try {
            const res = await this.createAxios().get('/files')
            return res.data
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            } else {
                throw new Error("Failed to get all saved files.");
            }
        }
    }

    getFileById = async (id: string) => {
        try {
            const res = await this.createAxios().get(`/files/${id}`)
            return res.data
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            } else {
                throw new Error(`Failed to get file (file id: ${id}).`);
            }
        }
    }

    uploadFile = async (data: FormData) => {
        try {
            const res = await this.createAxios("multipart/form-data").post("/upload", data);
            return res.data
        } catch (error) {
            if (axios.isAxiosError(error)) {
                throw error;
            } else {
                throw new Error("Failed to upload new file.");
            }
        }
    };
}

export default Api
