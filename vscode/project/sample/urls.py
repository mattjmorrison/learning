from django.urls import path
from sample.views.person import Person

urlpatterns = [
    path('people/', Person.as_view())
]
