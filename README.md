# NGO Hope Website

A professional, scalable website for an NGO to manage its public presence and organize community events (like eye check-up camps).

## 🌟 Features
- **Public Frontend**: SEO-optimized pages for Home, About, Contact, and Donations.
- **Admin Panel**: Dynamic dashboard to manage upcoming events and the image gallery.
- **Modern Tech Stack**: Built with Go (Golang), PostgreSQL, and Tailwind CSS.
- **Infrastructure**: Fully containerized with Docker for easy one-command deployment.
- **Quality Guard**: Integrated `golangci-lint` for strict code quality and data validation.

## 🛠️ Tech Stack
- **Backend**: Go 1.23
- **Database**: PostgreSQL 16
- **Frontend**: Go HTML Templates + Tailwind CSS
- **Infrastructure**: Docker & Docker Compose
- **QA**: golangci-lint, go-playground/validator

## 🚀 How to Run

### Prerequisites
- [Docker](https://www.docker.com/get-started) installed on your machine.

### Quick Start
1. **Clone the repository**:
   ```bash
   git clone https://github.com/ravi271088/ngo-site.git
   cd ngo-site
   ```

2. **Launch the application**:
   ```bash
   docker-compose -f deployments/docker-compose.yml up --build
   ```

3. **Access the site**:
   - **Website**: [http://localhost:8080](http://localhost:8080)
   - **Health Check**: [http://localhost:8080/health](http://localhost:8080/health)

### Admin Panel Usage
The admin panel allows you to manage events and images. 
- **Authentication**: The API currently uses a secret token for authorization: `secret-admin-token-123`.
- **Manage Events**: Add, edit, or delete medical camps and community activities.
- **Manage Gallery**: Upload image URLs and captions to showcase NGO work.

## 📁 Project Structure
- `cmd/server/`: Application entry point.
- `internal/`: Core business logic (Handlers, Services, Repository).
- `web/`: Frontend templates and static assets.
- `deployments/`: Docker configurations.
- `.golangci.yml`: Linter configuration for the "Guardian" system.

## 🛡️ Quality Assurance
To run the linter locally:
```bash
golangci-lint run
```
