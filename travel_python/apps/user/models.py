from django.db import models
from django.contrib.auth.models import AbstractUser


# Create your models here.


class User(AbstractUser):
    avatar = models.CharField(null=True, blank=True, verbose_name="头像")
    age = models.IntegerField(null=True, blank=True, verbose_name="年龄")

    class Meta:
        db_table = "tbl_user"


class Dynamic(models.Model):
    user = models.ForeignKey(User, on_delete=models.SET_NULL)
    create_time = models.DateTimeField(auto_now=True)

    class Meta:
        db_table = "tbl_user_dynamic"
