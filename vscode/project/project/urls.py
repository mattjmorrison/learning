from django.contrib import admin
from django.urls import path, include
from sample import urls as sample_urls

urlpatterns = [
    path('admin/', admin.site.urls),
    path('sample/', include(sample_urls))
]
