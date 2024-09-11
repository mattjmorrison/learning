from django import http

class Person:

    @classmethod
    def as_view(cls):
        def view(*args, **kwargs):
            return http.HttpResponse("[{}]")
        return view