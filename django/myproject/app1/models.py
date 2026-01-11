from django.db import models

class Item(models.Model):
    name = models.CharField(max_length=100)
    description = models.TextField()
    text = models.TextField()

    def __str__(self):
        return self.name

# python manage.py makemigrations
# python manage.py migrate
# python manage.py createsuperuser