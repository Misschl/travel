from django.db import models
from django.contrib.auth.models import AbstractUser


# Create your models here.


class User(models.Model):
    email = models.EmailField()
    password = models.CharField(max_length=50)
    username = models.CharField(max_length=18, null=True, blank=True)
    SEX_CHOICES = (
        (1, "男"),
        (0, "女"),
    )
    avatar = models.CharField(null=True, blank=True, verbose_name="头像", max_length=528)
    age = models.IntegerField(null=True, blank=True, verbose_name="年龄")
    sex = models.SmallIntegerField(choices=SEX_CHOICES, null=True, blank=True)
    introduce = models.TextField(null=True, blank=True)

    create_time = models.DateTimeField(auto_now=True)
    last_login_time = models.DateTimeField(auto_now_add=True, null=True, blank=True)
    is_delete = models.BooleanField(default=False)
    active = models.BooleanField(default=False)

    class Meta:
        db_table = "tbl_user"

# class Dynamic(models.Model):
#     user = models.ForeignKey(User, on_delete=models.SET_NULL, null=True)
#     create_time = models.DateTimeField(auto_now=True)
#
#     class Meta:
#         db_table = "tbl_user_dynamic"
