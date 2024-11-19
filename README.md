# **[WIP] Omnichain Analytics Dashboard**

### **Description**
The Omnichain Analytics Dashboard is a powerful tool that provides real-time insights into wallet balances, transactions, and cross-chain activities across multiple blockchains. By leveraging the LayerZero protocol, this dashboard seamlessly communicates across different blockchains, enabling a unified view of on-chain data.

### **Key Features**
- **Cross-Chain Wallet Tracking**: Monitor wallet balances across Ethereum, Binance Smart Chain, and other LayerZero-supported blockchains.
- **Transaction History**: Access recent transaction data from multiple chains in one place.
- **LayerZero Integration**: Powered by LayerZero, the dashboard facilitates secure and efficient omnichain communication.
- **User-Friendly Interface**: A simple, clean interface designed to make blockchain analytics accessible to everyone.

### **Technologies Used**
- **Backend**: Golang (for high-performance API development).
- **Frontend**: React (for a responsive and interactive UI).
- **Blockchain Interaction**: 
  - LayerZero SDK for omnichain communication.
  - `go-ethereum` for Ethereum-compatible chains.
- **HTTP Client**: Resty (for LayerZero API and blockchain queries).
- **Frameworks**:
  - Gin (for backend REST API).
  - Material-UI (for frontend components).

### **How LayerZero is Used**
LayerZero is the backbone of the dashboard’s cross-chain functionality, enabling:
1. Real-time balance synchronization across chains.
2. Secure and efficient communication between blockchain ecosystems.
3. Scalability to integrate additional blockchains with minimal overhead.

### **Installation and Setup**

#### **Prerequisites**
- [Go](https://golang.org/) (version 1.20+)
- [Git](https://git-scm.com/)
- Access to LayerZero endpoints and supported blockchain nodes (e.g., Infura).

#### **Steps to Run Locally**
1. **Clone the Repository**
   ```bash
   git clone https://github.com/your-username/omnichain-analytics-dashboard.git
   cd omnichain-analytics-dashboard
   ```

2. **Backend Setup**
   - Install dependencies:
     ```bash
     go mod tidy
     ```
   - Run the backend server:
     ```bash
     go run cmd/main.go
     ```
     
3. **Environment Configuration**
   - Create a `.env` file with the following keys:
     ```env
     INFURA_PROJECT_ID=your_infura_project_id
     LAYERZERO_API_URL=https://layerzero.api.endpoint
     ```

4. **Access the Dashboard**
   Open your browser and go to `http://localhost:3000` (or the appropriate port).

### **Future Enhancements**
- Add support for more blockchains (e.g., Avalanche, Solana).
- Implement gas fee analysis and cost optimization suggestions.
- Integrate NFT tracking for cross-chain assets.
- Provide real-time notifications for wallet activity.
