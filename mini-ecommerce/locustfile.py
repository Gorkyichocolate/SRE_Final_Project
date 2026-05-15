from locust import HttpUser, task

class EcommerceUser(HttpUser):

    @task
    def heavy(self):
        self.client.get("/heavy")